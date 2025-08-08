# Changelog

All notable changes to systemd-persistence-scanner are documented here.

### [2025-12-22]
- refactor: simplify token parsing pipeline and reduce cognitive complexity

### [2025-12-28]
- fix: resolve race condition during concurrent worker initialization

### [2026-01-18]
- test: add unit tests for boundary input cases and error branches

### [2026-02-12]
- feat: add support for custom timeout configuration via CLI flags

### [2026-02-17]
- security: sanitize input strings to mitigate format string risks

### [2026-02-21]
- style: format code according to style conventions

### [2026-03-14]
- refactor: use enum types for status codes instead of magic numbers

### [2026-04-28]
- security: sanitize input strings to mitigate format string risks

### [2026-05-21]
- test: add fuzzing harness for packet decoding routine

### [2026-05-28]
- fix: handle nil pointer dereference on unexpected connection close

### [2026-06-13]
- refactor: extract validation logic into dedicated helper module

### [2026-06-18]
- docs: update license headers and author metadata

### [2026-06-20]
- perf: minimize redundant heap allocations in hot loop

### [2026-07-04]
- perf: replace linear search with hash map lookup for fast querying

### [2026-07-07]
- feat: implement verbose output mode for troubleshooting

### [2026-08-08]
- docs: clarify prerequisite installation steps in README

### [2026-08-09]
- refactor: simplify token parsing pipeline and reduce cognitive complexity

### [2026-08-12]
- test: add unit tests for boundary input cases and error branches

### [2026-08-27]
- refactor: decouple configuration loader from runtime engine

### [2026-08-28]
- style: format code according to style conventions

