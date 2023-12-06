#!/usr/bin/env bash
sum=0
part=2
echo "Calculating..."
if [[ $part -eq 1 ]]; then
  redMax=12
  greenMax=13
  blueMax=14
  while read -r line; do
    gi=$(echo $line | cut -d":" -f1 | cut -d" " -f2)
    valid="yes"

    # TODO: can use positive lookahead
    while read -r redStr; do
      redCount=$(echo $redStr | cut -d" " -f1)
      if [[ $redCount -gt $redMax ]]; then valid="no"; fi
    done < <(echo $line | grep -oE "[0-9]+ red")

    while read -r greenStr; do
      greenCount=$(echo $greenStr | cut -d" " -f1)
      if [[ $greenCount -gt $greenMax ]]; then valid="no"; fi
    done < <(echo $line | grep -oE "[0-9]+ green")

    while read -r blueStr; do
      blueCount=$(echo $blueStr | cut -d" " -f1)
      if [[ $blueCount -gt $blueMax ]]; then valid="no"; fi
    done < <(echo $line | grep -oE "[0-9]+ blue")

    if [[ "$valid" = "yes" ]]; then sum=$(($sum + $gi)); fi
  done < ./input
else
  while read -r line; do
    gi=$(echo $line | cut -d":" -f1 | cut -d" " -f2)

    redMax=-1
    while read -r redStr; do
      redCount=$(echo $redStr | cut -d" " -f1)
      if [[ $redCount -gt $redMax ]]; then redMax="$redCount"; fi
    done < <(echo $line | grep -oE "[0-9]+ red")

    greenMax=-1
    while read -r greenStr; do
      greenCount=$(echo $greenStr | cut -d" " -f1)
      if [[ $greenCount -gt $greenMax ]]; then greenMax="$greenCount"; fi
    done < <(echo $line | grep -oE "[0-9]+ green")

    blueMax=-1
    while read -r blueStr; do
      blueCount=$(echo $blueStr | cut -d" " -f1)
      if [[ $blueCount -gt $blueMax ]]; then blueMax="$blueCount"; fi
    done < <(echo $line | grep -oE "[0-9]+ blue")

    power=$(($redMax * $greenMax * $blueMax))

    sum=$((sum + power))
  done < ./input
fi

echo "Total sum: $sum"
