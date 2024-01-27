package org.monkey.lang

import com.intellij.extapi.psi.PsiFileBase
import com.intellij.psi.FileViewProvider

open class MonkeyFile(fileViewProvider: FileViewProvider) : PsiFileBase(fileViewProvider, MonkeyLanguage) {
    override fun getFileType() = MonkeyFileType
    override fun toString() = "Monkey File"
}
