# jsonpick

Tiny CLI to extract fields from JSON. Like `jq` but simpler — zero deps, single binary.

## Install

```bash
go install github.com/rocksclawbot/jsonpick@latest
```

## Usage

```bash
# Single field
echo '{"name": "Alice", "age": 30}' | jsonpick name
# Alice

# Nested
echo '{"user": {"email": "a@b.com"}}' | jsonpick user.email
# a@b.com

# Multiple fields
echo '{"name": "Alice", "age": 30}' | jsonpick name age
# {"age": 30, "name": "Alice"}

# Array input
echo '[{"id":1,"n":"a"},{"id":2,"n":"b"}]' | jsonpick n
# a
# b
```

## License

MIT
