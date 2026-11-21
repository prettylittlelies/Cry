#include <dave/dave_interfaces.h>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

using namespace discord::dave;

static std::vector<uint8_t> unhex(const std::string& s) {
    std::vector<uint8_t> out;
    for (size_t i = 0; i + 1 < s.size(); i += 2) {
        out.push_back(static_cast<uint8_t>(std::stoi(s.substr(i, 2), nullptr, 16)));
    }
    return out;
}

static std::string hex(const std::vector<uint8_t>& b) {
    static const char* digits = "0123456789abcdef";
    std::string out;
    for (auto c : b) {
        out.push_back(digits[c >> 4]);
        out.push_back(digits[c & 15]);
    }
    return out;
}

int main() {
    std::ios::sync_with_stdio(false);
    std::unique_ptr<mls::ISession> session;
    std::unique_ptr<IEncryptor> encryptor = CreateEncryptor();
    std::map<std::string, std::unique_ptr<IDecryptor>> decryptors;
    std::string self;
    std::string line;
    while (std::getline(std::cin, line)) {
        std::istringstream in(line);
        std::string cmd, a, b;
        in >> cmd >> a >> b;
        std::string reply = "ok";
        if (cmd == "init") {
            self = a;
            session = mls::CreateSession(nullptr, "", [](std::string const& s, std::string const& r) {
                std::cerr << "mls failure " << s << ": " << r << std::endl;
            });
            std::shared_ptr<::mlspp::SignaturePrivateKey> key;
            session->Init(1, std::stoull(b), self, key);
        } else if (cmd == "recognize") {
            session->AddRecognizedUser(a);
        } else if (cmd == "extsender") {
            session->SetExternalSender(unhex(a));
        } else if (cmd == "keypackage") {
            reply = "ok " + hex(session->GetMarshalledKeyPackage());
        } else if (cmd == "proposals") {
            auto out = session->ProcessProposals(unhex(a));
            reply = out ? "ok " + hex(*out) : "ok";
        } else if (cmd == "commit") {
            auto r = session->ProcessCommit(unhex(a));
            reply = std::holds_alternative<RosterMap>(r) ? "ok" : std::holds_alternative<ignored_t>(r) ? "err ignored" : "err failed";
        } else if (cmd == "welcome") {
            reply = session->ProcessWelcome(unhex(a)) ? "ok" : "err failed";
        } else if (cmd == "auth") {
            reply = "ok " + hex(session->GetLastEpochAuthenticator());
        } else if (cmd == "encrypt") {
            encryptor->SetKeyRatchet(session->GetKeyRatchet(self));
            encryptor->AssignSsrcToCodec(1, Codec::Opus);
            auto frame = unhex(a);
            std::vector<uint8_t> out(encryptor->GetMaxCiphertextByteSize(MediaType::Audio, frame.size()));
            size_t written = 0;
            auto rc = encryptor->Encrypt(MediaType::Audio, 1, MakeArrayView<const uint8_t>(frame.data(), frame.size()), MakeArrayView(out), &written);
            out.resize(written);
            reply = rc == IEncryptor::Success ? "ok " + hex(out) : "err encrypt " + std::to_string(rc);
        } else if (cmd == "decrypt") {
            auto& dec = decryptors[a];
            if (!dec) {
                dec = CreateDecryptor();
            }
            dec->TransitionToKeyRatchet(session->GetKeyRatchet(a));
            auto frame = unhex(b);
            std::vector<uint8_t> out(dec->GetMaxPlaintextByteSize(MediaType::Audio, frame.size()));
            size_t written = 0;
            auto rc = dec->Decrypt(MediaType::Audio, MakeArrayView<const uint8_t>(frame.data(), frame.size()), MakeArrayView(out), &written);
            out.resize(written);
            reply = rc == IDecryptor::Success ? "ok " + hex(out) : "err decrypt " + std::to_string(rc);
        } else {
            reply = "err unknown";
        }
        std::cout << reply << std::endl;
    }
}
