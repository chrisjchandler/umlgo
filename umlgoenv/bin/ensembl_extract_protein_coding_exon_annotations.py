#!/home/cchandler/umlgo/umlgoenv/bin/python3
# -*- coding: utf-8 -*-
import re
import sys
from genometools.ensembl.extract_protein_coding_exon_annotations import main
if __name__ == '__main__':
    sys.argv[0] = re.sub(r'(-script\.pyw|\.exe)?$', '', sys.argv[0])
    sys.exit(main())
