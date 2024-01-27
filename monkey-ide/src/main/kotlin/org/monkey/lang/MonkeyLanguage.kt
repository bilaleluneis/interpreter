package org.monkey.lang

import com.intellij.lang.Language

object MonkeyLanguage : Language("monkey") {
    private fun readResolve(): Any = MonkeyLanguage
    override fun isCaseSensitive() = true
    override fun getDisplayName() = "M"
}
