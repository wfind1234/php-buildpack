<?php

# Create Redis connection
$redis = new Redis();
$redis->connect('127.0.0.1', 6379);

# Check if connection is intact
$redis->ping();
