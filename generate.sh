for size in 32 48 64 72 96 128 144 192 256 512
do
    convert -geometry ${size}x${size} -filter box aoisensi.svg aoisensi${size}px.png
done
ln aoisensi32px.png aoisensi.png

convert aoisensi.svg -define icon:auto-resize -filter box aoisensi.ico