#!/bin/sh
for file in $1*$2.ts
 do
    echo mv "$file" "${file//$2.ts/$2.model.ts}";
    mv "$file" "${file//$2.ts/$2.model.ts}";
 done