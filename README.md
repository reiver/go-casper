# go-casper

Package casper provides tools for working with the **Casper** file data format, for the Go programming language.

Typical uses for Casper is for configuration files, for API responses, for logs, for data transmissions, for serialization, etc.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-casper

[![GoDoc](https://godoc.org/github.com/reiver/go-casper?status.svg)](https://godoc.org/github.com/reiver/go-casper)


## Casper File Data Format

The **Casper** (CAscading PropERties) _file data format_ is inspired by the CSS (Cascading Style Sheets) file data format.
In that it looks a lot like it.
Although there are some differences.
For example:
```
apple {
	color:	red;
	type:	fruit;
}

banana {
	color:	yellow;
	type:	berry;
}

cherry {
	color: purple;
	type:	fruit;
}
```

But (instead of being used for _style sheets_, as with CSS) **Casper** is meant to be used used a
human-readble, human-editable, and machine-readable file data format, to be used for serializing
**hierarchical key-value pairs**.

**Casper** emphasizes software developer user experience (UX).

Typical uses for **Casper** is for configuration files, for API responses, for logs, for data transmissions, for serialization, etc.

**Casper** is meant to be a more human friendly alternative to JSON, YAML, XML, etc.

In particular, **Casper** is meant to be an improvement, in terms of developer user experience (UX), over JSON by:

* supporting comments,
* allowing trailing key-value pair deliminators,
* supporting the merging of multiple files into one, and
* supporting cascading,
* allowing for a notation that allows both the key, and value to be sliced out.


## Sample Casper

Here is what a simple **Casper** file might look like:
```
apple {
	color:	red;
	type:	fruit;
}

banana {
	color:	yellow;
	type:	berry;
}

cherry {
	color: purple;
	type:	fruit;
}
```

Here is a sample **Casper** file used as a config file:
```
database {
	username: "joeblow";
	password: "password123";
}

http {
	port: 8080;

	log {
		dest: stdout;
	}
}

telnet/port: 2323;
telnet/log/dest: stdout;
```

Here is the same thing written 3 different ways:
```
http {
	port: 8080;

	log {
		dest: stdout;
	}
}
```

```
http {
	port: 8080;
}

http/log {
		dest: stdout;
}
```

```
http/port: 8080;

http/log/dest: stdout;
```
