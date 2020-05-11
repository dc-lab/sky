#!/bin/bash

echo "$SKY_RM_TOKEN" >> /token

exec ./agent
