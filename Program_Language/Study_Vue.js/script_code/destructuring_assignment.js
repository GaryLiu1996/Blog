let obj = {p: ['hello', {y: 'world'}]}
let {p, p: [x, {y: z}]} = obj
console.log("--", p, x, z)

let {length: len} = 'hellow'//类似数组的对象中都有一个length属性
console.log("---", len)

let {t = 3} = {t: undefined}
let {v = 3} = {v: null}
let {b = 3} = {b: 4}
console.log("--", t, v, b)

// let {prop:a}= undefined
// let {prop:c}= null//TypeError
let {toString: cc} = 123
let {toString: aa} = true
console.log("--", cc === Number.prototype.toString, aa === Boolean.prototype.toString())
