object Hamming {
  def compute(s1: String, s2: String): Int = {
    require(s1.length == s2.length)
    return s1.zip(s2).count(x => x._1 != x._2)
  }
}
