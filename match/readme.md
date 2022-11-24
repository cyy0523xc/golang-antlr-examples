# Match

```bash
alias antlr4='java -Xmx500M -cp "../../antlr-4.11.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'

java -jar ../antlr-4.11.1-complete.jar -Dlanguage=Go -o parser Match.g4

java -jar ../antlr-4.11.1-complete.jar -o java_parser Match.g4
cd java_parser
javac -classpath ../../antlr-4.11.1-complete.jar *.java
java -jar ../../antlr-4.11.1-complete.jar org.antlr.v4.gui.TestRig Match prog ../samples/s01.txt -tokens
java -jar ../../antlr-4.11.1-complete.jar org.antlr.v4.gui.TestRig Match prog ../samples/s01.txt -gui
```


