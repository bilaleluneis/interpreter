package org.monkey.lang

import com.intellij.openapi.fileTypes.LanguageFileType

object MonkeyFileType : LanguageFileType(MonkeyLanguage) {
    override fun getName() = "Monkey File"
    override fun getDescription() = "Monkey lang file"
    override fun getDefaultExtension() = "mky"
    override fun getIcon() = MonkeyIcons.FILE
}
