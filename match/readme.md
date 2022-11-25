# Match

```bash
pip install antlr4-tools

# 生成go代码
antlr4 -Dlanguage=Go -o parser Match.g4

antlr4

antlr4-parse Match.g4 prog -tree samples/s07.txt
antlr4-parse Match.g4 prog -gui samples/s07.txt
antlr4-parse Match.g4 prog -tokens samples/s07.txt
```


