#!/bin/sh
echo 'Runing migrations...'
/demo-services-simple_app/bin/migrate up > /dev/null 2>&1 &

echo 'Start application...'
/demo-services-simple_app/bin/app