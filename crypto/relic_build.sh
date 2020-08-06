#!/bin/bash

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

pushd "$DIR"

# Ensure the directory is writeable
chmod -R +w "$(pwd)"

# Clone RELIC if needed
if [ ! -d "relic" ]; then
git clone --depth 1 --branch relic-toolkit-0.5.0 git@github.com:relic-toolkit/relic.git
fi

mkdir -p "$DIR/relic/build"
pushd "$DIR/relic/build"

# Set RELIC config for Flow

#
GENERAL=(-DTIMER=CYCLE -DCHECK=OFF -DVERBS=OFF)
LIBS=(-DSHLIB=OFF -DSTLIB=ON)
COMP=(-DCOMP="-O3 -funroll-loops -fomit-frame-pointer -march=native -mtune=native")
RAND=(-DRAND=HASHD -DSEED=)

#
BN_REP=(-DALLOC=AUTO -DALIGN=1 -DWSIZE=64 -DBN_PRECI=1024 -DBN_MAGNI=DOUBLE)
ARITH=(-DARITH=EASY)  # try GMP or X64-ASM-382
PRIME=(-DFP_PRIME=381)

#
BN_METH=(-DBN_KARAT=0 -DBN_METHD="COMBA;COMBA;MONTY;SLIDE;STEIN;STRON")
FP_METH=(-DFP_KARAT=0 -DFP_METHD="INTEG;INTEG;INTEG;MONTY;LOWER;SLIDE")
PRIMES=(-DFP_PMERS=OFF -DFP_QNRES=ON -DFP_WIDTH=2)
FPX_METH=(-DFPX_METHD="INTEG;INTEG;LAZYR")
CURVES=(-DEP_MIXED=ON -DEP_PLAIN=OFF -DEP_SUPER=OFF -DEP_DEPTH=4 -DEP_WIDTH=2)
PP_METH=(-DPP_METHD="LAZYR;OATEP")

# Generate make files using cmake
cmake "${COMP[@]}" "${GENERAL[@]}" \
    "${LIBS[@]}" "${RAND[@]}" \
    "${BN_REP[@]}" "${ARITH[@]}" \
    "${PRIME[@]}" "${PRIMES[@]}" \
    "${CURVES[@]}" \
    "${BN_METH[@]}" \
    "${FP_METH[@]}" \
    "${FPX_METH[@]}" \
    "${PP_METH[@]}" ..

# Compile the static library
make clean
make relic_s -j8
rm -f CMakeCache.txt

popd
popd