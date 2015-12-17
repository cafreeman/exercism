class Phrase(val p: String) {
  def wordCount: Map[String, Int] = {
    return p.split(' ')
      .groupBy(_.distinct)
      .map(x => x._1 -> x._2.length)
  }
}
