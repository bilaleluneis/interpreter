package org.monkey.lang.psi


import com.intellij.psi.tree.IElementType
import org.monkey.lang.MonkeyLanguage

open class MonkeyTokenType(debugName: String) : IElementType(debugName, MonkeyLanguage) {
    override fun toString() = "MonkeyTokenType." + super.toString()
}
