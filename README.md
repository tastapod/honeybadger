# A ratel is a honey badger, and honey badger don't care

Read any structured data without knowing its schema in advance.

Useful for transforming or extracting values from large or unknown documents to e.g. transform one set of YAML to another.

Go's YAML and JSON parsing assumes that you know in advance the data you are unmarshalling, and expects you to provide appropriate structs to load the data into.

I want something closer to Python's YAML and JSON modules that generate dictionaries and lists containing values. Since Go is a statically typed language, we cannot just go stuffing arbitrary values into maps, and in particular we cannot mix different types in the same map or list.

My approach is to create an `Any` type that can store any value. See [design ideas](design-ideas.md) for different approaches we could try.
