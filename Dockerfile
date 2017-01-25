# from scratch in order to generate the smallest image possible.
FROM scratch
# Add and run the previously built GOlang static binary instead of building it on the image.
ADD service /
CMD ["/service"]
