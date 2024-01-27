package org.monkey.lang

import com.intellij.testFramework.ParsingTestCase
import java.io.IOException

internal abstract class MonkeyParserTestBase : ParsingTestCase("", "mky", MonkeyParserDef()) {
    override fun getTestDataPath() = "src/test/resources"

    override fun skipSpaces() = true

    override fun doTest(checkErrors: Boolean) {
        doTestImpl(checkResult = true, ensureNoErrorElements = true)
        if (checkErrors) {
            assertFalse(
                "PsiFile contains error elements",
                toParseTreeText(myFile, skipSpaces(), includeRanges()).contains("PsiErrorElement")
            )
        }
    }

    private fun doTestImpl(checkResult: Boolean, ensureNoErrorElements: Boolean) {
        val name = testName.trim()
        try {
            parseFile(name, loadFile("$name.$myFileExt"))
            if (checkResult) {
                checkResult(name, myFile)
                if (ensureNoErrorElements) {
                    ensureNoErrorElements()
                }
            } else {
                toParseTreeText(myFile, skipSpaces(), includeRanges())
            }
        } catch (e: IOException) {
            throw RuntimeException(e)
        }
    }
}
