{\rtf1\ansi\ansicpg1252\cocoartf1504\cocoasubrtf830
{\fonttbl\f0\fnil\fcharset0 Monaco;\f1\fmodern\fcharset0 Courier;}
{\colortbl;\red255\green255\blue255;\red0\green0\blue0;\red255\green255\blue255;}
{\*\expandedcolortbl;;\csgenericrgb\c0\c0\c0;\csgenericrgb\c100000\c100000\c100000;}
\margl1440\margr1440\vieww10800\viewh8400\viewkind0
\deftab720
\pard\pardeftab720\sl360\partightenfactor0

\f0\fs24 \cf2 \cb3 \expnd0\expndtw0\kerning0
FROM mongo\
COPY customers.json /customers.json\
CMD mongoimport --host mongodb --db exampleDb --collection contacts --type json --file /customers.json --jsonArray
\f1\fs30 \
}