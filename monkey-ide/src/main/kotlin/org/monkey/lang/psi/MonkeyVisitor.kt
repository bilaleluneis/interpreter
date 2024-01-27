package org.monkey.lang.psi

import com.intellij.psi.PsiElement
import com.intellij.psi.PsiElementVisitor

class MonkeyVisitor : PsiElementVisitor() {
    fun visitProperty(prorp: MonkeyProperty) = visitPsiElement(prorp)
    private fun visitPsiElement(element: PsiElement) = visitElement(element)
}
