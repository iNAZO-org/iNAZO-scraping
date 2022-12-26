/*
  [input]
  pos: number

  [output]
  string (:===:で要素を区切り連結している)

  [description]
  Seleinumで実行される。北大の成績分布の一覧ページで、一行の要素を「:===:」で連結して返す。
*/
{
  document.getElementsByXPath = function (expression, parentElement) {
    var r = [];
    var x = document.evaluate(
      expression,
      parentElement || document,
      null,
      XPathResult.ORDERED_NODE_SNAPSHOT_TYPE,
      null
    );
    for (var i = 0, l = x.snapshotLength; i < l; i++) {
      r.push(x.snapshotItem(i));
    }
    return r;
  };
  return document
    .getElementsByXPath(
      '//*[@id="gvResult"]/tbody/tr[count(td)=18][position()=' + pos + "]/td"
    )
    .map((v) => v.textContent)
    .join(":---:");
}
