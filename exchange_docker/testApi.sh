curl -X POST -H "Content-Type: application/json" localhost:8080/exchange/latest -d '
{
  "baseCurrency": "EUR",
  "targetCurrency": "NOK"
}
'
echo
