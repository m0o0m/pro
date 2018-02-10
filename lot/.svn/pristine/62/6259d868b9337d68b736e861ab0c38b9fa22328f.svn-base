/**
 * Created by huygo on 2017/12/6.
 */
// 传入数组和需要删除的对象
const splice_arr = {
  splice_arr(_arr, _obj){
    var length = _arr.length;
    for (var i = 0; i < length; i++) {
      if (_arr[i] == _obj) {
        if (i == 0) {
          _arr.shift(); //删除并返回数组的第一个元素
          return;
        }
        else if (i == length - 1) {
          _arr.pop();  //删除并返回数组的最后一个元素
          return;
        }
        else {
          _arr.splice(i, 1); //删除下标为i的元素
          return;
        }
      }
    }
  }
}
export default splice_arr;
