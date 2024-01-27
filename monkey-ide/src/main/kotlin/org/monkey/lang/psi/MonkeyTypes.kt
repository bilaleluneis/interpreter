package org.monkey.lang.psi

import com.intellij.lang.ASTNode
import com.intellij.psi.PsiElement
import com.intellij.psi.tree.IElementType
import org.monkey.lang.psi.impl.MonkeyPropertyImpl

// FIXME: remove @JvmField once MonkeyParser is translated to kotlin
object MonkeyTypes {
    @JvmField val PROPERTY: IElementType = MonkeyElementType("PROPERTY")
    @JvmField val COMMENT: IElementType = MonkeyTokenType("COMMENT")
    @JvmField val CRLF: IElementType = MonkeyTokenType("CRLF")
    @JvmField val KEY: IElementType = MonkeyTokenType("KEY")
    @JvmField val SEPARATOR: IElementType = MonkeyTokenType("SEPARATOR")
    @JvmField val VALUE: IElementType = MonkeyTokenType("VALUE")

    fun createElement(node: ASTNode): PsiElement {
        val type = node.elementType
        if (type === PROPERTY) {
            return MonkeyPropertyImpl(node)
        }
        throw AssertionError("Unknown element type: $type")
    }
}

