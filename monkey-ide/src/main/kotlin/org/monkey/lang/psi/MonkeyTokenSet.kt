package org.monkey.lang.psi

import com.intellij.psi.tree.TokenSet

object MonkeyTokenSet {
    val KEY = TokenSet.create(MonkeyTypes.KEY)
    val VALUE = TokenSet.create(MonkeyTypes.VALUE)
    val SEPARATOR = TokenSet.create(MonkeyTypes.SEPARATOR)
    val COMMENTS = TokenSet.create(MonkeyTypes.COMMENT)
}
