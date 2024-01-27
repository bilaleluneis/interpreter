package org.monkey.lang.parser

import com.intellij.lang.ASTNode
import com.intellij.lang.PsiBuilder
import com.intellij.lang.PsiParser
import com.intellij.lang.parser.GeneratedParserUtilBase.*
import com.intellij.psi.tree.IElementType
import org.monkey.lang.psi.MonkeyTypes.PROPERTY

class MonkeyParser : PsiParser {

    override fun parse(elementType: IElementType, psiBuilder: PsiBuilder): ASTNode = eofParser(elementType, psiBuilder)

    private fun eofParser(rootElement: IElementType, psiBuilder: PsiBuilder): ASTNode {
        val builder = adapt_builder_(rootElement, psiBuilder, this, null)
        val rootMarker = enter_section_(builder, 0, _NONE_, PROPERTY, "<property>")
        while (!builder.eof()) {
            val tokType = builder.tokenType
            builder.advanceLexer()
        }
        rootMarker.done(rootElement)
        val ast = builder.treeBuilt
        return ast
    }

    private fun nParser(elementType: IElementType, psiBuilder: PsiBuilder): ASTNode {
        val builder = adapt_builder_(elementType, psiBuilder, this, null)
        val marker = enter_section_(builder, 0, _COLLAPSE_, null)
        //TODO: parse root look at MonkyParser.kt parseRoot: boolean = parse(elemenType, builder)
        exit_section_(builder, 0, marker, elementType, true /*from parse root*/, true, TRUE_CONDITION)
        val ast = builder.treeBuilt
        return ast
    }

}
