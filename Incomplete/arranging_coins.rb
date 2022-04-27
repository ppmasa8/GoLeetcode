# @param {Integer} n
# @return {Integer}
def arrange_coins(n)
    num = 0
    for i in 1 .. 2 ** 31
        num += i
        return i - 1 if num > n
    end
end
