"""Parser DDL sik.sql -> metadata kolom."""
import re

import os

SQL = os.environ.get('SIK_SQL', os.path.join(os.path.dirname(os.path.abspath(__file__)), '..', '..', '..', 'source', 'sik.sql'))
_cache = None


def tables():
    global _cache
    if _cache is not None:
        return _cache
    _cache = {}
    cur = None
    for line in open(SQL, encoding='latin-1'):
        m = re.match(r'^CREATE TABLE `(\w+)`', line)
        if m:
            cur = {'name': m.group(1), 'cols': [], 'pk': [], 'fks': {}}
            _cache[cur['name']] = cur
            continue
        if cur is None:
            continue
        if line.startswith(')'):
            cur = None
            continue
        s = line.strip().rstrip(',')
        m = re.match(r"^`(\w+)` (\w+)(\(((?:[^()']|'[^']*')*)\))?(.*)$", s)
        if m:
            rest = m.group(5)
            cur['cols'].append({
                'name': m.group(1), 'type': m.group(2).lower(), 'arg': m.group(4),
                'notnull': 'NOT NULL' in rest, 'autoinc': 'AUTO_INCREMENT' in rest,
            })
            continue
        m = re.match(r'^PRIMARY KEY \((.*)\)', s)
        if m:
            cur['pk'] = re.findall(r'`(\w+)`', m.group(1))
            continue
        m = re.match(r'^CONSTRAINT `\w+` FOREIGN KEY \(`(\w+)`\) REFERENCES `(\w+)` \(`(\w+)`\)', s)
        if m:
            cur['fks'][m.group(1)] = (m.group(2), m.group(3))
    return _cache


def camel(name):
    out = ''.join(p[:1].upper() + p[1:] for p in re.split(r'[_\s]+', name) if p)
    if out and out[0].isdigit():
        out = 'X' + out
    return out


INTS = {'int', 'bigint', 'smallint', 'tinyint', 'mediumint', 'year'}
FLOATS = {'double', 'float', 'decimal', 'real'}


def enum_values(c):
    return re.findall(r"'([^']*)'", c['arg'] or '')


def gormtype(c):
    return c['type'] + (f"({c['arg']})" if c['arg'] else '')
