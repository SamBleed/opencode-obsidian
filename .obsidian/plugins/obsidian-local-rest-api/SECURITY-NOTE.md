# Obsidian Local REST API — Local Configuration Policy

`data.json` may contain API keys, certificates and private keys. It must stay local and must not be committed or exported in shared ZIPs.

Use `data.example.json` as a safe reference. Keep the insecure HTTP server disabled unless you explicitly need it in an isolated local environment.

Required defaults for Bunker OS:

- `enableInsecureServer`: `false`
- real API keys: local only
- private keys: local only
- shared artifacts: never include `data.json`
