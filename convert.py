import copy
import json
import os
import re
import string
import sys

raw_books = [(x, open(os.path.join('books', x), 'r').read()) for x in os.listdir('./books')]

header = re.compile('\*\*\*\s?START.*$', re.M)
footer = re.compile('\*\*\*\s?END.*$', re.M)
metadata = {
  'title': re.compile('Title:\s?(.*)$', re.M),
  'author': re.compile('Author:\s?(.*)$', re.M)
}

for fname, book in raw_books:
  data = {}
  for n, r in metadata.iteritems():
    result = r.search(book)
    if result:
      data[n] = result.group(1).strip()

  data['id'] = int(fname[2:-4])
  content = footer.split(header.split(book)[1])[0]
  data['data'] = []
  for p in content.strip().split('\r\n\r\n'):
    if (len(p) == 0):
      continue

    p = string.replace(p.strip(), '\r\n', " ")
    paragraph = {
      'text': p,
      'characters': len(p),
      'sentences': len(p.split('.')),
    }
    data['data'].append(paragraph)

  data['paragraphs'] = len(data['data'])

  print '%s: %s' % (data['id'], data['paragraphs'])

  fobj = open("gae/books/%s.json" % (data['id'],), 'wb')
  fobj.write(json.dumps(data))
  fobj.close()
