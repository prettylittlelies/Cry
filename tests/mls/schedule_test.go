package mls_test

import (
	"testing"

	"github.com/prettylittlelies/cry/mls"
)

func TestKeySchedule(t *testing.T) {
	type epoch struct {
		CommitSecret   hexBytes `json:"commit_secret"`
		PSKSecret      hexBytes `json:"psk_secret"`
		GroupContext   hexBytes `json:"group_context"`
		Joiner         hexBytes `json:"joiner_secret"`
		Welcome        hexBytes `json:"welcome_secret"`
		Init           hexBytes `json:"init_secret"`
		SenderData     hexBytes `json:"sender_data_secret"`
		Encryption     hexBytes `json:"encryption_secret"`
		Exporter       hexBytes `json:"exporter_secret"`
		EpochAuth      hexBytes `json:"epoch_authenticator"`
		External       hexBytes `json:"external_secret"`
		Confirmation   hexBytes `json:"confirmation_key"`
		Membership     hexBytes `json:"membership_key"`
		Resumption     hexBytes `json:"resumption_psk"`
		ExternalPub    hexBytes `json:"external_pub"`
		ExporterVector struct {
			Label   string
			Context hexBytes
			Length  int
			Secret  hexBytes
		} `json:"exporter"`
	}
	type vector struct {
		InitialInit hexBytes `json:"initial_init_secret"`
		Epochs      []epoch
	}
	s := suite(t)
	for _, v := range load[vector](t, "key-schedule.json") {
		init := []byte(v.InitialInit)
		for i, e := range v.Epochs {
			var ctx mls.GroupContext
			if err := mls.Unmarshal(e.GroupContext, &ctx); err != nil {
				t.Fatalf("epoch %d: %v", i, err)
			}
			joiner := s.JoinerSecret(init, e.CommitSecret, &ctx)
			equal(t, "joiner", joiner, e.Joiner)
			equal(t, "welcome", s.WelcomeSecret(joiner, e.PSKSecret), e.Welcome)
			sec := s.EpochSecrets(joiner, e.PSKSecret, &ctx)
			for _, c := range []struct {
				name      string
				got, want []byte
			}{
				{"init", sec.Init, e.Init}, {"sender data", sec.SenderData, e.SenderData},
				{"encryption", sec.Encryption, e.Encryption}, {"exporter", sec.Exporter, e.Exporter},
				{"authenticator", sec.Authenticator, e.EpochAuth}, {"external", sec.External, e.External},
				{"confirmation", sec.Confirmation, e.Confirmation}, {"membership", sec.Membership, e.Membership},
				{"resumption", sec.Resumption, e.Resumption},
			} {
				equal(t, c.name, c.got, c.want)
			}
			ext, err := s.ExternalKey(sec.External)
			if err != nil {
				t.Fatal(err)
			}
			equal(t, "external pub", ext.PublicKey().Bytes(), e.ExternalPub)
			x := e.ExporterVector
			equal(t, "export", s.Export(sec.Exporter, x.Label, x.Context, x.Length), x.Secret)
			init = sec.Init
		}
	}
}

func TestTranscriptHashes(t *testing.T) {
	type vector struct {
		ConfirmationKey hexBytes `json:"confirmation_key"`
		Content         hexBytes `json:"authenticated_content"`
		InterimBefore   hexBytes `json:"interim_transcript_hash_before"`
		ConfirmedAfter  hexBytes `json:"confirmed_transcript_hash_after"`
		InterimAfter    hexBytes `json:"interim_transcript_hash_after"`
	}
	s := suite(t)
	for _, v := range load[vector](t, "transcript-hashes.json") {
		var content mls.AuthenticatedContent
		if err := mls.Unmarshal(v.Content, &content); err != nil {
			t.Fatal(err)
		}
		confirmed := s.ConfirmedTranscriptHash(v.InterimBefore, &content)
		equal(t, "confirmed", confirmed, v.ConfirmedAfter)
		equal(t, "confirmation tag", s.MAC(v.ConfirmationKey, confirmed), content.Auth.ConfirmationTag)
		equal(t, "interim", s.InterimTranscriptHash(confirmed, content.Auth.ConfirmationTag), v.InterimAfter)
	}
}
