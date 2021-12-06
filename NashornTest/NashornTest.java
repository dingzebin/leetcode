package leetcode.algorithms;

import jdk.nashorn.api.scripting.NashornScriptEngine;

import javax.script.ScriptEngineManager;
import javax.script.ScriptException;

/**
 * @author dzeb
 * @version 1.0
 * @Description TODO
 * @createTime 2021/10/31 22:39
 */
public class NashornTest {
    public static void main(String[] args) throws ScriptException {
        NashornScriptEngine engine = (NashornScriptEngine) new ScriptEngineManager().getEngineByName("nashorn");
        System.out.println(engine.compile("a();"));
    }
}
