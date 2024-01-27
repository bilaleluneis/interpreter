package org.monkey.lang

import com.intellij.lang.ASTNode
import com.intellij.lang.ParserDefinition
import com.intellij.lang.PsiParser
import com.intellij.lexer.Lexer
import com.intellij.openapi.project.Project
import com.intellij.psi.FileViewProvider
import com.intellij.psi.PsiElement
import com.intellij.psi.PsiFile
import com.intellij.psi.tree.IFileElementType
import com.intellij.psi.tree.TokenSet
import org.monkey.lang.parser.MonkeyParser
import org.monkey.lang.psi.MonkeyTokenSet
import org.monkey.lang.psi.MonkeyTypes

class MonkeyParserDef : ParserDefinition {
    override fun createLexer(project: Project): Lexer = MonkeyLexer()
    override fun createParser(project: Project): PsiParser = MonkeyParser()
    override fun getFileNodeType(): IFileElementType = IFileElementType(MonkeyLanguage)
    override fun getCommentTokens(): TokenSet = MonkeyTokenSet.COMMENTS
    override fun getStringLiteralElements(): TokenSet = TokenSet.EMPTY
    override fun createElement(astNode: ASTNode): PsiElement = MonkeyTypes.createElement(astNode)
    override fun createFile(fileViewProvider: FileViewProvider): PsiFile = MonkeyFile(fileViewProvider)
}
