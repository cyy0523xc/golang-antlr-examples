// Makefile for debugging Match.g4
antlr4_exe = java -jar /usr/local/lib/antlr-4.11.1-complete.jar
grun_exe = java org.antlr.v4.gui.TestRig
target =

all:
    go build

gen:
    $(antlr4_exe) -Dlanguage=Go -o parser Match.g4

gen_visitor:
    $(antlr4_exe) -Dlanguage=Go -visitor -o parser Match.g4

gen_java:
    $(antlr4_exe) Match.g4

gui: gen_java
    javac *.java
    $(grun_exe) Match prog $(target) -gui

trace: gen_java
    javac *.java
    $(grun_exe) Match prog $(target) -trace

tokens: gen_java
    javac *.java
    $(grun_exe) Match prog $(target) -tokens

tree: gen_java
    javac *.java
    $(grun_exe) Match prog $(target) -tree
