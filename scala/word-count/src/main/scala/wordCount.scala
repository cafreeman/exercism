class Phrase(val p: String) {
  def wordCount: Map[String, Int] = {
    return p.split(Array(' ', ','))
      .map(
        _.filter(x => {
          x.isLetterOrDigit || x == '\''
        }).toLowerCase
      )
      .groupBy(x => x)
      .filter(x => x._1.nonEmpty)
      .map(x => x._1 -> x._2.length)
  }
}
