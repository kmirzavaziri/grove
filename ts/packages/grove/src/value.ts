export type Value = number | string | boolean | Struct | List;
export type Struct = {[key: string]: Value};
export type List = Array<Value>;
