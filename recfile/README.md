# Recutils

A Go library to access the human-editable, plain text databases called GNU recfiles

# Known Limitations

According to the GNU recutils specification: "Comments should be treated the same way as other data in the database". However, this library discards them

# Examples

For now, the tests are actually simple enough that they should serve as good examples of how to use the core library functions (i.e. `Load()`, `Save()`). Please see `database_test.go` :-)
