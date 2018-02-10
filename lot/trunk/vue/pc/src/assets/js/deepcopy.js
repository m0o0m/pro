const mydeepCopy = function deepCopy(o, c) {
  var c = c || {}
  for (var i in o) {
    if (typeof o[i] === 'object') { //要考虑深复制问题了
      if (o[i].constructor === Array) {
        //这是数组
        c[i] = []
      } else {
        //这是对象
        c[i] = {}
      }
      deepCopy(o[i], c[i])
    } else {
      c[i] = o[i]
    }
  }
  return c
}

export default mydeepCopy;


