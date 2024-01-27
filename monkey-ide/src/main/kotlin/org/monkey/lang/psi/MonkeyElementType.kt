package org.monkey.lang.psi

import com.intellij.psi.tree.IElementType
import org.monkey.lang.MonkeyLanguage

open class MonkeyElementType(debugName: String) : IElementType(debugName, MonkeyLanguage)
