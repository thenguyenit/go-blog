We are building a website which sales lottery around the world, many issues we met such as languages, networking, security, performance, database migration...and before launch may be OK.
But unfortunately, the big problem is the daylight saving time issue or often called DST occurred.

So the questions:

* What is Daylight saving time?
* What is the best practice for timezone?

In a nice day on March, when Americans and Europeans change their clocks, our website should follow it, but unfortunately, this does not occur. 

And our problem is storing all DateTime with UTC timezone to MySQL and using addHour in Datetime class PHP and DATE_ADD in MySQL with an offset value in the range (-12 to +12) for displaying to end user.

**Totally wrong.**

**What is daylight saving time?**

**Daylight saving time (DST)** is the practice of setting the clock forward one hour from the standard time during the summer months and back again in the fall, in order to make better use of natural daylight.

Most of the United States begins DST at 2:00 am on the second Sunday in March back again on the first Sunday in November.

In the European Union, DST begins at 1:00 am on the last Sunday in March and ends the last Sunday in October.

**What is the best practice for timezone?**

1. The time according to a unified standard that is not affected by DST (UTC or GMT).

2. Don't store offset, should be store timezone(ex: America/Porto_Acre).

3. The timezone offsets are not always and an integer of hours ( Indian Standard Time is UTC+05:30, and Nepal uses UTC+05:45).

4. In PHP, should use the native time zone conversions provided by DateTime class.

5. In javascript, should use moment.js.

6. In MySql, should use the native timezone conversions provided by CONVERT_TZ function. ([Detail](https://dev.mysql.com/doc/refman/5.5/en/time-zone-upgrades.html))

7. Keep OS, database and application timezone in sync and common are UTC.

8. Never trust the timestamp you get from client machines as valid.

9. When testing, make sure you test countries in Western, Eastern, Northern, Southern and countries have DST and not use DST. Test boundary cases, such as a timezone is UTC+12 or UTC+13 or haft-hour timezone, at least.

Exceptions:

- In some cases, you may need to store both the UTC time and the equivalent local time.