##
#  League-a-lot
#
# @file
# @version 0.1

#all: league-a-lot

.PHONY: docker image deploy

docker: image
	docker push schuermannator/league-arm:latest

image:
	docker build -t schuermannator/league-arm .

# end
