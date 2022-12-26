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
      '//*[@id="gvResult"]/tbody/tr[position()=' + pos + "]/td"
    )
    .map((v) => v.textContent)
    .join(":---:");
}
