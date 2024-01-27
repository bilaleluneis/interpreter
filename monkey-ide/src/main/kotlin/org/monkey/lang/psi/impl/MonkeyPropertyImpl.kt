package org.monkey.lang.psi.impl

import com.intellij.extapi.psi.ASTWrapperPsiElement
import com.intellij.lang.ASTNode
import com.intellij.psi.PsiElementVisitor
import org.monkey.lang.psi.MonkeyProperty
import org.monkey.lang.psi.MonkeyVisitor

class MonkeyPropertyImpl(node: ASTNode) : ASTWrapperPsiElement(node), MonkeyProperty {
    private fun accept(visitor: MonkeyVisitor) = visitor.visitProperty(this)
    override fun accept(visitor: PsiElementVisitor) = if (visitor is MonkeyVisitor) {
        accept(visitor)
    } else super.accept(visitor)
}
