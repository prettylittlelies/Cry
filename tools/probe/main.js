const { app, BrowserWindow, session } = require('electron');
const fs = require('fs');

const [scenarioPath, lab] = process.argv.slice(-2);
const scenario = JSON.parse(fs.readFileSync(scenarioPath, 'utf8'));
const hosts = scenario.hosts.map((h) => `MAP ${h} ${lab}`).join(',');

app.commandLine.appendSwitch('host-resolver-rules', hosts);
app.commandLine.appendSwitch('ignore-certificate-errors');
app.commandLine.appendSwitch('disable-features', 'LocalNetworkAccessChecks,PrivateNetworkAccessSendPreflights,PrivateNetworkAccessRespectPreflightResults,BlockInsecurePrivateNetworkRequests');

setTimeout(() => { console.error('timeout'); app.exit(1); }, 60000);

app.whenReady().then(async () => {
  try {
    for (const c of scenario.cookies ?? []) await session.defaultSession.cookies.set(c);
    const w = new BrowserWindow({ show: false });
    await w.loadURL(scenario.page);
    await w.webContents.executeJavaScript(`(${scenario.script})()`);
    app.exit(0);
  } catch (e) {
    console.error(e);
    app.exit(1);
  }
});
