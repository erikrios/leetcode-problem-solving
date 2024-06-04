class Solution {
  fun compressedString(word: String): String {
    val sb = StringBuilder()
    var previousChar = word.first()
    var count = 1
    for (i in 1 until word.length) {
        val currentChar = word[i]
        if (previousChar == currentChar) {
            count++
        } else {
            appendCompressed(sb, previousChar, count)
            previousChar = currentChar
            count = 1
        }
    }
    appendCompressed(sb, previousChar, count)
    return sb.toString()
  }

  fun appendCompressed(sb: StringBuilder, char: Char, count: Int) {
    var remainingCount = count
    while (remainingCount > 9) {
        sb.append(9).append(char)
        remainingCount -= 9
    }
    if (remainingCount > 0) {
        sb.append(remainingCount).append(char)
    }
  }
}
