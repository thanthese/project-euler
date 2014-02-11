US_currency = [100, 50, 25, 10, 5, 1]
US_target_amount = 100

english_currency = [200, 100, 50, 20, 10, 5, 2, 1]
english_target_amount = 200

coin = english_currency
target_amount = english_target_amount


def recur(total_so_far, coin_index):
    coin_value = coin[coin_index]
    if total_so_far == target_amount:
        return 1
    if coin_value == 1:
        return 1
    will_fit = (target_amount - total_so_far) / coin_value
    return sum(recur(total_so_far + coin_value * n, coin_index + 1)
               for n in range(will_fit + 1))


print recur(0, 0)
