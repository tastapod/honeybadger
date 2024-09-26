# Design ideas

## The `Document` type

The top level will be a `Document` that contains (or represents?) either a `map[string]Any` or a `[]Any`.

## The `Any` interface

`Any` is an interface which will be implemented by `AnyYaml` and `AnyJson` types. The interface will emerge through tests. Methods will tell you what the Any stores and access the value under various conversions. Most accessors will return an error-value pair.

 An `Any` is either a scalar, a list, or a map. Perhaps these last two are both aggregates, where an aggregate is simply anything not a scalar.

You can decode a scalar value using a `Decode` method that takes a reference to a scalar variable. We will lean on e.g. the `yaml.3` `Node.Decode` method that does some introspection of the variable to determine how to interpret the value.

You can iterate over a list/slice or over a map's key-value pairs.

## The `AnyYaml` type

`AnyYaml` stores the `yaml3.Node`  and some metadata about what the value represents
