<p align="center"><img src="docs/assets/logo.svg" width="88" alt="My Search logo"></p>

# My Search

Find authorised information across your connected Workspace apps. An original Apache-2.0 One Workspace application; not an upstream fork.

![My Search application](docs/assets/screenshot.jpg)

## Windows download

Download the Windows x64 ZIP from [Releases](https://github.com/oneworkspacestudio/my-search/releases). Verify SHA256SUMS.txt, extract to a program folder and run **MySearch.exe**. No Docker, WSL, Node.js or Python is needed at runtime. A modern browser and a dedicated **NTFS** data folder are required. This is a **0.1.0-preview.1** release; unsigned Windows executables may prompt for trust.

The simplest managed installation is [One Workspace](https://github.com/oneworkspacestudio/one-workspace). New default apps are installed during manager updates unless explicitly excluded. Existing app data and user files are preserved.

## Included

Live, permission-aware federation for Drive file metadata (including Docs/Photos), Notes, Calendar, Contacts and Tasks; source status and partial failures; per-source pagination where supported; app/type/date/tag/owner filters where metadata exists; safe highlighting, authorised suggestions and bounded single-edit spelling retry; opt-in personal recent searches; global Workspace shortcut.

## Standalone

```powershell
.\MySearch.exe --data C:\WorkspaceData\my-search --listen 127.0.0.1:8795
```

The app opens a one-time first-owner setup link. Choose your own password. Later starts reuse the same data folder. Keep private data separate from the executable and user-created files. Use a distinct port for each app. App Settings offers encrypted backup and account details; managed installations use MyAccounts automatically.

## Important limits

My Vault is always excluded. Mail bodies and private Chat messages are not searched; unsupported apps expose launch metadata only. No global administrator crawler or foreign database access. Drive searches at most 20 allowed sources/100 results and requires refinement rather than fake pagination. Calendar/Contacts use bounded collection APIs; their links open the app/date, not a record-specific selection. Spelling retry searches a prefix only after no exact result, then accepts single-edit title matches; it is not a universal fuzzy index. Results refresh while visible every 15 seconds; clicking still invokes source authorisation. Filters exclude records lacking the selected metadata. Standalone mode has no remote sources until hosted with Workspace.

See [operations and recovery](docs/OPERATIONS.md), [API contracts](docs/API.md), [acceptance evidence](docs/ACCEPTANCE.md), [architecture](docs/ARCHITECTURE.md) and [security](SECURITY.md). Preview status does not imply clean-machine, reboot, Windows-service or independent penetration-test certification.

## Build locally

Install PowerShell 7, Go 1.27.2 and Node.js 24+ for development only. Run `scripts/build.ps1 -Go <go.exe> -Npm <npm.cmd> -Release` from a clean committed checkout. The script uses pinned Go/npm dependencies, regenerates `.aui` UI, runs TypeScript, Go tests/vet and frontend contracts, builds the Windows icon/EXE, inventories dependency licenses and emits an immutable ZIP plus checksums. Browser tests: set APP_E2E_URL, APP_E2E_PASSWORD and run `npm run test:e2e` in web against disposable data.

No GitHub Actions are used. Distribution includes LICENSE, NOTICE, full dependency license texts, dependency inventory, compatibility and source build information. Source edits belong in `web/src/screens/*.aui`; generated React is reviewed and committed.
