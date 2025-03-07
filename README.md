# Grove

## Overview
TODO

## Gex
TODO

## Executor
TODO

## Flux
TODO

## Guard

TODO

## GuardX

TODO

### Error Template

Each Guard takes some templates in its config to return in failed guard scenarios. The error template is
nothing but a string that can contain variables.

`value ${value} should be at least ${config.min}`

The following variables are always available.

- `value`. The value that is being guardd.
- `config`. The guard config.

Additional variables might be available, depending on the guard and error scenario. You can find additional
variables in each error scenario in the following tables.

| Guard           | Config                            | Error Templates             | Type Requirement  | Description                                                                                  |
|-----------------|-----------------------------------|-----------------------------|-------------------|----------------------------------------------------------------------------------------------|
| **Common**      |                                   |                             |                   |                                                                                              |
| `Type`          | `TypeName` (ripd.Type)            | Err                         | -                 | Ensure the value is of the specified type.                                                   |
| `Empty`         | -                                 | Err                         | -                 | Ensure value is nil or empty (`""`, `[]`, `{}`, or `nil`), but doesn't allow `false` or `0`. |
| `Not`           | `Guard` (Guard)                   | Err                         | -                 | Not of the underlying guard.                                                                 |
| `And`           | `Guards` ([]Guard)                | Err (`nested_err`)          | -                 | And of the underlying guard.                                                                 |
| `Or`            | `Guards` ([]Guard)                | Err                         | -                 | Or of the underlying guard.                                                                  |
| **Collections** |                                   |                             |                   |                                                                                              |
| `Len`           | `Min` (uint64), `Max` (uint64)    | ErrType, ErrMin, ErrMax     | map, list, string | Check that the value length is within specified `min` and `max` inclusive range.             |
| `All`           | `Guard` (Guard)                   | ErrType, Err (`nested_err`) | list, map         | Ensure all nested values pass the specified guard.                                           |
| `Any`           | `Guard` (Guard)                   | ErrType, Err                | list, map         | Ensure at least one nested value passes the specified guard.                                 |
| `RequiredKeys`  | `Keys` ([]string)                 | ErrType, Err (`key`)        | map               | Ensure specified keys exist in the map.                                                      |
| `AllKeys`       | `Pattern` (string)                | ErrType, Err                | map               | Ensure all keys match the provided regex pattern.                                            |
| `AnyKey`        | `Pattern` (string)                | ErrType, Err                | map               | Ensure at least one key matches the provided regex pattern.                                  |
| `NestMap`       | `Key` (string), `Guard` (Guard)   | ErrType, Err (`nested_err`) | map               | Guard that `value[key]` passes the specified guard.                                          |
| `NestList`      | `Index` (uint64), `Guard` (Guard) | ErrType, Err (`nested_err`) | list              | Guard that `value[index]` passes the specified guard.                                        |
| `UniqueItems`   | -                                 | ErrType, Err (`item`)       | list              | Ensure all items in the list are unique.                                                     |
| **Numbers**     |                                   |                             |                   |                                                                                              |
| `Range`         | `Min` (number), `Max` (number)    | ErrType, ErrMin, ErrMax     | number            | Ensure value is within the specified inclusive range.                                        |
| `Integer`       | -                                 | ErrType, Err                | number            | Ensure the value is an integer.                                                              |
| **Strings**     |                                   |                             |                   |                                                                                              |
| `Match`         | `Pattern` (string)                | ErrType, Err                | string            | Ensure the value matches the specified regex pattern.                                        |
| `Contains`      | `Substring` (string)              | ErrType, Err                | string            | Ensure the value contains the specified substring.                                           |
| `HasPrefix`     | `Prefix` (string)                 | ErrType, Err                | string            | Ensure the value starts with the specified prefix.                                           |
| `HasSuffix`     | `Suffix` (string)                 | ErrType, Err                | string            | Ensure the value ends with the specified suffix.                                             |
| `Enum`          | `Values` ([]string)               | ErrType, Err                | string            | Ensure value is one of the accepted values.                                                  |
| `Email`         | -                                 | ErrType, Err                | string            | Guard that the value is in proper email format.                                              |
| `Phone`         | -                                 | ErrType, Err                | string            | Guard that the value is in proper phone format.                                              |
| **Bools**       |                                   |                             |                   |                                                                                              |
| `Equal`         | `Value` (bool)                    | ErrType, Err                | bool              | Ensure the value is equal to the specified boolean.                                          |

### Writing Your Own Custom Guards
TODO

## Grove
TODO

TODO Error Component
TODO Fragment Component

## GroveX
TODO

| Component     | TODO | Args | Props | Description                                                                                                                                                           |
|---------------|------|------|-------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Displays**  |      |      |       |                                                                                                                                                                       |
| DBreadCrumb   |      | TODO | TODO  | TODO                                                                                                                                                                  |
| DButton       |      | TODO | TODO  | TODO                                                                                                                                                                  |
| DDivider      |      | TODO | TODO  | TODO                                                                                                                                                                  |
| DImage        | TODO | TODO | TODO  | TODO                                                                                                                                                                  |
| DProfile      |      | TODO | TODO  | TODO                                                                                                                                                                  |
| DTypography   |      | TODO | TODO  | TODO                                                                                                                                                                  |
| **Inputs**    |      |      |       |                                                                                                                                                                       |
| IBool         | TODO | TODO | TODO  | TODO handles checkbox and switch                                                                                                                                      |
| IDatetime     | TODO | TODO | TODO  | TODO handles date, time, and datetime                                                                                                                                 |
| IList         | TODO | TODO | TODO  | TODO handles a list of group of inputs (similar to django admin tabular / inline) TODO better name?                                                                   |
| IMultiSelect  | TODO | TODO | TODO  | TODO handles multiselect, single select, typeahead (autocomplete), with static source or dynamic source, can render as radio, checkbox set, and toggle button as well |
| INumber       | TODO | TODO | TODO  | TODO handles number, can have prefix and suffix (for example kg will be suffixed for weight)                                                                          |
| IPhoneNumber  | TODO | TODO | TODO  | TODO handles the rendering of prefix of country code                                                                                                                  |
| ISingleSelect | TODO | TODO | TODO  | TODO handles multiselect, single select, typeahead (autocomplete), with static source or dynamic source, can render as radio, checkbox set, and toggle button as well |
| IText         | TODO | TODO | TODO  | TODO                                                                                                                                                                  |
| **Layouts**   |      |      |       |                                                                                                                                                                       |
| LAccordion    | TODO | TODO | TODO  | TODO                                                                                                                                                                  |
| LBox          | TODO | TODO | TODO  | TODO                                                                                                                                                                  |
| LGrid         | TODO | TODO | TODO  | TODO                                                                                                                                                                  |
| LMasonry      | TODO | TODO | TODO  | TODO                                                                                                                                                                  |
| LPage         |      | TODO | TODO  | TODO                                                                                                                                                                  |
| LTabs         | TODO | TODO | TODO  | TODO                                                                                                                                                                  |
| **Specials**  |      |      |       |                                                                                                                                                                       |
| XClickable    | TODO | TODO | TODO  | TODO                                                                                                                                                                  |
| XModal        | TODO | TODO | TODO  | TODO                                                                                                                                                                  |

### Writing Your Own Custom Components

TODO

# To Do List

1. Unmarshal `gex.Value` to any given concrete type, then use it in walk, when trying to unmarshal error node data
2. Should we use `gex.Value` and marshal / unmarshal it, or use `any` instead, with type assertions?
3. Should we merge Data and Execution in flux for simplicity?
4. Add logging to examples/blog server.
5. Add tests for guards: not, and, or
6. Fix the common guard `Type` marshalling procedure. (`func (v Type) Marshal() *gex.Value`)
7. Use match for email guard instead of new script?
8. Put data besides input and children when sending to fe in render? 
