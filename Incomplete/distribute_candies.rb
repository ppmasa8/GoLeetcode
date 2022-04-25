# @param {Integer[]} candy_type
# @return {Integer}
def distribute_candies(candy_type)
    capa = candy_type.size/2
    kind = candy_type.uniq.size
    return capa >= kind ? kind : capa
end
