# ADR004: My Search

Decision: original Apache-2.0 app at oneworkspacestudio/my-search. Start with
browser-federated per-user requests to documented, installed same-origin APIs:
Drive /api/tools/search, Notes /api/notes?q=, Calendar /api/events, Contacts
/api/contacts and My Tasks /api/tasks?q=. The real portal registry /_workspace/apps
is authoritative for availability; extend its app IDs/capabilities while retaining
old fields. Browser cookies remain scoped to their source apps. No shared admin
credentials, cross-app DB reads, global crawler or server-side snippets cache.
Sources requiring a first SSO connection are shown explicitly, not silently
searched under a different identity. Docs and Photos are Drive metadata results.
Unsupported sources have metadata-only launcher discovery; Mail bodies and Chat
are opt-in only when user APIs are verified. Vault is excluded even from source
selection and suggestions. Refresh current user-authorised results while visible;
clear on logout, source failure or account change. Results use text-only rendering,
per-source pagination, stable IDs, safe deep links, source status and deduplication.
Global header search and keyboard shortcut link to the same app, with graceful
absence when Search was unchecked. Recent terms are off by default and scoped to
immutable identity if enabled. No unsupported endpoint is guessed.

Shared decision: native Go server, embedded AUI-generated React/TypeScript UI,
SQLite (modernc, BSD-3-Clause), protected per-app NTFS data folders and one EXE.
Reuse the existing One Workspace Apache-2.0 platform, OIDC, branding and supervisor
conventions through a public pkg/apphost package in the actual suite repository.
Canonicalise the suite Go module path to its existing organization repository;
there are no supported external consumers of its internal packages. Each original
app has a dedicated organization repository, rather than a misleading fork.
OIDC uses code+PKCE, state cookie, nonce, issuer/audience verification and immutable
issuer/subject keys. Session/CSRF secrets stay on the server. Local bootstrap is
loopback-only and explicitly disabled in managed mode. No GitHub Actions: user
instruction overrides the brief's Actions-template suggestion; run gates locally.
