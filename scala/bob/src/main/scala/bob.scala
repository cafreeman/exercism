
class Bob {
  def hey(words: String): String = {
    if (Bob.isSilent(words)) {
      return "Fine. Be that way!"
    } else if (Bob.isShouting(words) && !Bob.isOnlyNumbers(words)) {
      return "Whoa, chill out!"
    } else if (Bob.isQuestion(words)) {
      return "Sure."
    }
    return "Whatever."
  }
}

object Bob {
  def getLetters (words: String): String = {
    return words.filter(_.isLetter)
  }

  def isShouting(words: String): Boolean = {
    return getLetters(words).forall(_.isUpper)
  }

  def isQuestion(words: String): Boolean = {
    return words.last == '?'
  }

  def isOnlyNumbers(words: String): Boolean = {
    return words.filter(_.isLetterOrDigit).forall(_.isDigit)
  }

  def isSilent(words: String): Boolean = {
    return words.isEmpty || words.forall(_.isSpaceChar)
  }
}
