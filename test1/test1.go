package main

import (
	"fmt"
	"github.com/jjcinaz/mibtool/smi"
	"log"
)

func main() {
	mib := smi.NewMIB("/Users/jjc/.snmp/mibs", "/usr/local/share/snmp/mibs", "/usr/share/snmp/mibs")
	mib.Debug = true
	err := mib.LoadModules("RADIO-BRIDGE-MIB")
	if err != nil {
		log.Fatal(err)
	}

	/*	// Walk all symbols in MIB
		mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
			fmt.Printf("%-40s %s %s\n", sym, oid, sym.Description)
		})
	*/
	// Look up OID for an OID string
	oid, err := mib.OID("RADIO-BRIDGE-MIB::rfMinRssi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oid.String())

	oid = smi.OID{1, 3, 6, 1, 4, 1, 31926, 2, 2, 1, 18}
	fmt.Println(mib.SymbolString(oid))
}
