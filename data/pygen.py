import random
import string

seed = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

def generate_randomstr():
    strlen = random.randrange(8,16)
    sa = [random.choice(seed) for i in range(strlen)]
    return "".join(sa)

def main():
    res = []
    with open("testurl.txt", "w") as outputf:
        urlgen = []
        for _i in range(1000):
            rstr = generate_randomstr()
            rtime = random.randrange(1, 150)
            for _j in range(rtime):
                urlgen.append(rstr)
            res.append((rstr, rtime))
        random.shuffle(urlgen)
        for _, v in enumerate(urlgen):
            outputf.write(v + "\n")
    res = sorted(res, key=lambda a :a[1], reverse=True)
    with open("pygen_sorted_result.txt", "w") as outputf:
        for _, v in enumerate(res):
            outputf.write(v[0] + "======" + str(v[1]) + "\n")

if __name__ == "__main__":
    main()

