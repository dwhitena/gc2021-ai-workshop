#!/usr/bin/env bash
mkdir -p models
model_name='deepset/bert-base-cased-squad2'
./hf-importer --repo "models" --model "$model_name"
